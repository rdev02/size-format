package sizeformat

import (
	"strings"
	"testing"
)

func TestToString(t *testing.T) {
	runCompareToString(1, "1.00B", t)
	runCompareToString(10, "10.00B", t)
	runCompareToString(kilo, "1.00KB", t)
	runCompareToString(kilo+100, "1.10KB", t)
	runCompareToString(kilo+10, "1.01KB", t)
	runCompareToString(10*KB, "10.00KB", t)
	runCompareToString(10*MB+99*kilo, "10.10MB", t)
	runCompareToString(10*GB, "10.00GB", t)
	runCompareToString(5*TB, "5.00TB", t)
	runCompareToString(3*PB, "3.00PB", t)
	runCompareToString(306729213952, "306.73GB", t)
}

func runCompareToString(arg uint64, expected string, t *testing.T) {
	actual := ToString(arg)
	if strings.Compare(actual, expected) != 0 {
		t.Errorf("[%d] should be formatted to [%s], instead got [%s]", arg, expected, actual)
	}
}

func TestToInt(t *testing.T) {
	runCompareToInt("5 b", 5, t)
	runCompareToInt("1kb", kilo, t)
	runCompareToInt("1000 Mb", 1000*MB, t)
	runCompareToInt("1.2Gb", 1.2*GB, t)
	runCompareToInt("0.05 Tb", 0.05*TB, t)
	runCompareToInt(" 11.5PB", 11.5*PB, t)
}

func TestToIntFailures(t *testing.T) {
	runCompareToIntFailure("aTB", t)
	runCompareToIntFailure("1.2.3k", t)
	runCompareToIntFailure("1.3g", t)
	runCompareToIntFailure("-1.3Gb", t)
}

func runCompareToInt(arg string, expected uint64, t *testing.T) {
	actual, err := ToNum(&arg)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("[%s] should be translated to [%d], instead got [%d]", arg, expected, actual)
	}
}

func runCompareToIntFailure(arg string, t *testing.T) {
	actual, err := ToNum(&arg)
	if err == nil {
		t.Errorf("Error should have occurred, instead got result %d", actual)
	}
}
