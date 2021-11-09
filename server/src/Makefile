SHELL := /bin/bash
DEFAULT = help

PROJECTPATH = $(pwd)/
BUILDPATH = $(pwd)build/

DEBUGEXE = $(PROJECTPATH)

COMPILER = rustc
BUILDER = go

Q = @ 	# SILENCE the Bash commands ;)

all: $(DEFAULT)

help:
	$(Q)echo "--- WWWS Make"
	$(Q)echo "make run               - Runs debug executable"
	$(Q)echo "make build             - Builds executable"
	$(Q)echo "make clean-build       - Builds executable"
	$(Q)echo "make publish       	 - Prepares for packaging"

run:
	make build
	
	$(shell rm -dfr $(BUILDPATH) )
	
	@echo "Running debug version from build folder"
	$(shell mkdir -p $(BUILDPATH))

	# copy executable to build dir
	$(Q)$(shell mv src/onlinedtp $(BUILDPATH)) 	
	$(Q)$(shell cp -r src/templates/ $(BUILDPATH))

	# cd into it and run scaffold init 
	$(shell cd $(BUILDPATH) && ./onlinedtp) 
	
build:
	@echo "Building using $(BUILDER)..."
	$(Q)$(shell cd src/  && $(BUILDER) build)
	@echo "Moving binary into build directory"

	$(Q)$(shell mkdir -p $(BUILDPATH))
	$(Q)$(shell mv src/onlinedtp $(BUILDPATH))

clean-build:
	@echo "Cleaning build folder... $(BUILDPATH)"
	$(shell rm -rf $(BUILDPATH)* )

publish:
	@make build
	@echo "Publishing files to $(pwd)$(BUILDPATH)"
	@$(shell cp -r src/templates/ $(BUILDPATH))
clean:
	$(Q)$(BUILDER) clean
	$(shell rm -dfr $(BUILDPATH) )

test: 
	@echo $(PROJECTPATH)
