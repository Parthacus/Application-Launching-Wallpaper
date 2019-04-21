#!/bin/sh

terminal_wm_class="Surf"
terminal_exec="surf http://localhost:8081"

# no terminal started, so start one
if [ -z "`wmctrl -lx | grep Surf`" ]; then
    $terminal_exec &
else
    wmctrl -r '@cgDISVMf:- | Startpage' -e 0,0,3,1366,790
    wmctrl -x -a $terminal_wm_class
fi;
