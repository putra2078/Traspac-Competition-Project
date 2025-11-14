package presence

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"hrm-app/internal/domain/employee"
	"hrm-app/internal/domain/work_hour"
	// "hrm-app/internal/pkg/database"
	// "hrm-app/internal/pkg/utils"
)

type UseCase interface {
	// CheckIn(checkin *Presence) error
	Checkin(userID uint, lat float64, long float64) error
	Checkout(userID uint, lat float64, long float64) error
	// CheckOut(checkout *Presence) error
	// GetByDateRange(startDate, endDate time.Time) error
	// GetByEmployeeDateRange(employeeID *employee.Employee, startDate, endDate time.Time)
	// DeleteByID(id uint) error
}

type PresenceUseCase struct {
	repo         Repository
	employeeRepo employee.Repository
	workHourRepo work_hour.Repository
}

func NewUseCase(repo Repository, eRepo employee.Repository, wRepo work_hour.Repository) UseCase {
	return &PresenceUseCase{
		repo:         repo,
		employeeRepo: eRepo,
		workHourRepo: wRepo,
	}
}

func (u *PresenceUseCase) Checkin(userID uint, lat float64, long float64) error {
	employee, err := u.employeeRepo.FindByUserID(userID)
	if err != nil {
		return fmt.Errorf("failed to find employee: %w", err)
	}

	workHour, err := u.workHourRepo.FindByID(employee.WorkTime)
	if err != nil {
		return fmt.Errorf("failed to find work hour: %w", err)
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	// tanggal saja
	dateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	// jam checkin untuk kolom TIME (harus tanpa tanggal)
	checkInTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), now.Second(), 0, loc).Format("15:04:05")

	workStart, err := parseWorkTime(workHour.StartTime)
	if err != nil {
		return fmt.Errorf("failed to parse work start time: %w", err)
	}

	workStartToday := time.Date(now.Year(), now.Month(), now.Day(),
		workStart.Hour(), workStart.Minute(), workStart.Second(), 0, loc)

	checkInStatus := "On Time"
	if now.Before(workStartToday) {
		checkInStatus = "Too Early"
	} else if now.After(workStartToday.Add(10 * time.Minute)) {
		checkInStatus = "Late"
	}

	existing, err := u.repo.FindCheckinToday(employee.ID)
	if err == nil && existing != nil {
		return errors.New("kamu sudah melakukan check-in hari ini")
	}

	presence := &Presence{
		EmployeeID:    employee.ID,
		Date:          dateOnly,
		CheckInHour:   checkInTime,
		LatCheckIn:    lat,
		LongCheckIn:   long,
		CheckInStatus: checkInStatus,
	}

	return u.repo.CreateCheckin(presence)
}

func parseWorkTime(s string) (time.Time, error) {
	layouts := []string{
		"15:04:05-07:00", // ex: 08:00:00+07:00 or 08:00:00+00:00
		"15:04:05-07",    // ex: 08:00:00+07
		"15:04:05Z07:00", // alternative ISO-like
		"15:04:05",       // no offset
	}

	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, fmt.Errorf("empty time string")
	}

	// Normalisasi format offset
	if strings.HasSuffix(s, "+00") {
		s = strings.Replace(s, "+00", "+00:00", 1)
	}

	var lastErr error
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		} else {
			lastErr = err
		}
	}

	return time.Time{}, fmt.Errorf("failed to parse time %q: %v", s, lastErr)
}

func (u *PresenceUseCase) Checkout(userID uint, lat float64, long float64) error {
	// Cari data employee berdasarkan auth middleware
	employee, err := u.employeeRepo.FindByUserID(userID)
	if err != nil {
		return fmt.Errorf("failed to find employee: %w", err)
	}

	// Cari jam kerja employee
	workHour, err := u.workHourRepo.FindByID(employee.WorkTime)
	if err != nil {
		return fmt.Errorf("failed to find work hour: %w", err)
	}

	// Menentukan location untuk load waktu
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	checkOutTime := now.Format("15:04:05")
	// checkOutTime := time.Date(0, 1, 1, now.Hour(), now.Minute(), now.Second(), 0, loc)

	// Parsing
	workEnd, err := parseWorkTime(workHour.EndTime)
	if err != nil {
		return fmt.Errorf("failed to parse work end time: %w", err)
	}

	workEndToday := time.Date(
		now.Year(), now.Month(), now.Day(),
		workEnd.Hour(), workEnd.Minute(), workEnd.Second(),
		0, loc,
	)

	// Pengecekan check-out status
	checkOutStatus := "On Time"
	if now.Before(workEndToday) {
		checkOutStatus = "Too Early"
	} else if now.After(workEndToday.Add(10 * time.Minute)) {
		checkOutStatus = "Late"
	}

	uncheckout, err := u.repo.FindCheckOutToday(employee.ID)
	if err != nil {
		// return fmt.Errorf("failed to check uncheckout presence: %w", err)
		return fmt.Errorf("Anda sudah check-out hari ini")
	}

	if uncheckout == nil {
		return errors.New("tidak ada presensi yang perlu checkout")
	}

	// Update data presences
	uncheckout.CheckOutHour = &checkOutTime
	uncheckout.LatCheckOut = &lat
	uncheckout.LongCheckOut = &long
	uncheckout.CheckOutStatus = &checkOutStatus

	return u.repo.CheckOut(uncheckout)
}
