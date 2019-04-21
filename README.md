# Application-Launching-Wallpaper
An application launcher designed for openbox that works like a wallpaper but can also be accessed through the browser


Screenshot:
![Screenshot](https://raw.githubusercontent.com/Parthacus/Application-Launching-Wallpaper/master/screenshot1.png)

The launcher is hosted on localhost:8081 which can be accessed through a browser when the scripts are running, additionally, to make it function like an interactive wallpaper, you need to pick a browser (I went with surf browser for lightness), add the wm settings to hide it from pager to prevent it from showing up in the alt-tab menu, hide from taskbar so it doesn't show up in taskbar menus, to maximise it so it looks like a wallpaper, to hide the window decorations and layer it below all other applications by default (in the case of openbox):

```xml
    <application class="Surf">
      <decor>no</decor>
      <maximized>yes</maximized>
      <layer>below</layer>
      <desktop>all</desktop>
      <skip_pager>yes</skip_pager>
      <skip_taskbar>yes</skip_taskbar>
    </application>
  </applications>
  ```
  Then you can bind something like stickscript.sh to your Super + D key combo (before the minimise all windows action)
  stickscript.sh maximises surf, in my case, so when Super + D is pressed, the wallpaper is raised, creating the illusion it is the wallpaper. 
My openbox config for this looks like this:

```xml
    <keybind key="W-d">
      <action name="Execute">
        <command>bash ~/"stickscript.sh"</command>
      </action>
      <action name="ToggleShowDesktop"/>
    </keybind>
```
