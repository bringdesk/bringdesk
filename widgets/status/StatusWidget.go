package status

import "image/color"

type StatusWidget struct {
	title   string     /* Large status header with summary              */
	summary string     /* Summary information with value or status      */
	icon    string     /* Style icon provide user well known status bar */
	color   color.RGBA /* Base status color                             */

}

func Render() {
	/* Draw background */
	/* Draw icon */
	/* Draw title */
	/* Draw summary */
}
