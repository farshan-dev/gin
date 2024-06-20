## Using Make to Run Commands at a Glance
```# apt-get -y install make```

create Makefile

```# touch Makefile```

file structure:

```
command_name:
	command to run like cd ~

.PHONY: command_name

```

usage:

```# make command_name```

