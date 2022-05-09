/*
 * This file is part of Crocodile Game Bot.
 * Copyright (C) 2019  Viktor
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func logInit() {
	formatter := new(prefixed.TextFormatter)
	formatter.TimestampFormat = "2006/01/02 15:04:05"
	formatter.ForceFormatting = true
	formatter.ForceColors = true
	formatter.FullTimestamp = true
	formatter.DisableSorting = true

	log.Formatter = formatter
	log.SetOutput(os.Stdout)
}

func setLogLevel(level string) error {
	switch level {
	case "PANIC":
		log.SetLevel(logrus.PanicLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "WARN":
		log.SetLevel(logrus.WarnLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
	default:
		return errors.New("unknown log level")
	}
	return nil
}
