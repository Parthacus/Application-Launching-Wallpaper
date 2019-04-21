# Application-Launching-Wallpaper
An application launcher designed for openbox that works like a wallpaper but can also be accessed through the browser


Screenshot:
![Screenshot](https://raw.githubusercontent.com/Parthacus/Application-Launching-Wallpaper/master/2019-04-21-201734_1366x768_scrot.png)

The launcher is hosted on localhost:8081 which can be accessed through a browser when the scripts are running, additionally, to make it function like an interactive wallpaper, you need to pick a browser (I went with surf browser for lightness), add the wm settings to hide it from pager, taskbar, to maximise it, to hide the window decorations and layer it below all other applications by default (in the case of openbox):

```
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
  
  
