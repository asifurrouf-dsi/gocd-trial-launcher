package gocd

import (
	"github.com/gocd-contrib/gocd-trial-launcher/utils"
	"github.com/mgutz/ansi"
)

const LOGO_TEXT = `

    @@@
  @@@@@@@@
  /@@@@@@@@@@
    /@@@@@@@@@@@
       #@@@@@@@@@@@                      @@@@@@@@   @@@        @@@@@@@@@
          (@@@@@@@@@@@                @@@@@@@@@@@@@.@@@     @@@@@@@%@@@@@@.
             #@@@@@@@@@@@           ,@@@&         @@@@@    @@@@         @@@@
   @@@@@        /@@@@@@@@@@@        @@@            @@@@   @@@*           @@@@
  @@@@@@@          @@@@@@@@@@      @@@#             @@@  *@@@             @@@
  @@@@@@@       @@@@@@@@@@@@       @@@              @@@  @@@@             @@@
  @@@@@@@    @@@@@@@@@@@@          @@@@            @@@@   @@@             @@@
  @@@@@@@ @@@@@@@@@@@@              @@@@          @@@@@   @@@@           @@@,
  @@@@@@@@@@@@@@@@@                  @@@@@      @@@@@@@     @@@@      ,@@@@
  @@@@@@@@@@@@@@                       @@@@@@@@@@@  @@@       @@@@@@@@@@@
  @@@@@@@@@@@                                       @@@
  @@@@@@@@                                         @@@@
    @@@                             @@@@          @@@@
                                     @@@@@@@%%@@@@@@*
                                        @@@@@@@@@@

`

func PrintLogo() {
	utils.Out(ansi.Magenta + LOGO_TEXT + ansi.Reset)
}
