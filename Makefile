SHELL := /bin/bash
DEFAULT = help

PROJECTPATH = $(HOME)/Projekter/WebworkflowWithScribus
BUILDPATH = $(PROJECTPATH)/build/

DEBUGEXE = $(PROJECTPATH)

COMPILER = rustc
BUILDER = go

Q = @ 	# SILENCE the Bash commands ;)

all: $(DEFAULT)

help:
	$(Q)echo "--- rust-empty (0.7 005)"
	$(Q)echo "make run               - Runs debug executable"
	$(Q)echo "make build             - Builds executable"
	$(Q)echo "make clean-build       - Builds executable"

.PHONY: \
		run \
		build \
		clean \

run:
	make build
	
	$(shell rm -dfr $(BUILDPATH) )
	
	@echo "Running debug version from build folder"
	$(shell mkdir -p $(BUILDPATH))

	# copy executable to build dir
	$(shell mv ./onlinedtp $(BUILDPATH)) 
	
	# cd into it and run scaffold init 
	$(shell cd $(BUILDPATH) && ./onlinedtp) 
	
build:
	@echo "Building using $(BUILDER)..."
	$(Q)$(shell cd src/  && $(BUILDER) build)
	
	if [! -d ../build ]; then 
		$(shell mkdir -p ../build)
	fi
		$(shell mv ./onlinedtp ../build)

clean-build:
	@echo "Cleaning build folder... $(BUILDPATH)"
	$(shell rm -rf $(BUILDPATH)* )
clean:
	$(Q)$(BUILDER) clean
	$(shell rm -dfr $(BUILDPATH) )

test: 
	@echo $(PROJECTPATH)
