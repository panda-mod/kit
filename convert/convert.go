package convert

import "fmt"

func buildError(val any, target string) error {
	return fmt.Errorf("unable to convert %#v of type %T to %s", val, val, target)
}
