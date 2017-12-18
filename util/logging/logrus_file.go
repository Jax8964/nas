// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package logging

import (
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// EnableFileLogger enable log file output
func EnableFileLogger(path string) (err error) {
	if len(path) == 0 {
		// If the path is not set, the file log is not output
		return nil
	}
	if !filepath.IsAbs(path) {
		path, _ = filepath.Abs(path)
	}
	if err = os.MkdirAll(path, 0700); err != nil {
		return err
	}
	filePath := path + "/neb-%Y%m%d%H%M.log"
	linkPath := path + "/neb.log"
	writer, err := rotatelogs.New(
		filePath,
		rotatelogs.WithLinkName(linkPath),
		//rotatelogs.WithMaxAge(time.Duration(604800) * time.Second),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)

	if err != nil {
		return err
	}

	hook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.ErrorLevel: writer,
	}, nil)
	logrus.AddHook(hook)
	return nil

}
