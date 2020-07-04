# i3-tabn

Focus the Nth tab of the currently focused tabbed container in i3.

Useful for a workspace with a single tabbed container, e.g.

```
+------+------+------+------+
| tab0 | tab1 | .... | tabN |
+------+------+------+------+
|                           |
|                           |
|                           |
|                           |
+---------------------------+
```

## Install

`go get -u daniel.lipovetsky.me/i3-tabn`

Make sure that `$GOBIN` is in your path, or put `i3-tabn` somewhere in your path.

## Configure

Add key bindings to your i3 configuration. For example, if `$mod` is Meta, the following maps `Meta+N` to the `Nth` tab.

```
# Select Nth tab
bindsym $mod+1 exec --no-startup-id i3-tabn 0
bindsym $mod+2 exec --no-startup-id i3-tabn 1
bindsym $mod+3 exec --no-startup-id i3-tabn 2
bindsym $mod+4 exec --no-startup-id i3-tabn 3
bindsym $mod+5 exec --no-startup-id i3-tabn 4
```

