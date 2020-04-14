# phdis

Script to disable Pi-Hole for a certain amount of time. It just open an SSH connection to machine running the Pi-Hole and run the *disable* command.

Why not just use the web interface? Because this way is like 2 secs faster and i'm lazy.

## Usage

Change **username** and **ip_address** values in the first lines of the script, then:

```bash
./phdis.sh interval
```

e.g.: To stop Pi-Hole for 10 mi

```bash
./phdis.sh 10
```
