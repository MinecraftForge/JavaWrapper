NAME=ForgePreLauncher
MAJOR=0
MINOR=1
PATCH=4
VERSION=$(MAJOR).$(MINOR).$(PATCH)
BUILD=$(BUILD_NUMBER)
BIN_DIR=build/$(VERSION)
APP_DIR=$(BIN_DIR)/FPL.app

clean: 
	rm -rf build
bindir:
	mkdir -p $(BIN_DIR)

windows: 
	GOOS=windows GOARCH=386 go build -v -o $(NAME)-$(VERSION).$(BUILD).exe forgewrapper.go

	zip $(BIN_DIR)/FPL-Win.$(VERSION).$(BUILD).zip  $(NAME)-$(VERSION).$(BUILD).exe
	mv $(NAME)-$(VERSION).$(BUILD).exe $(BIN_DIR)/$(NAME)-$(VERSION).$(BUILD).exe
 
osx:
	cp -r .apptemplate $(APP_DIR)

	GOOS=darwin go build -v -o $(NAME)_OSX-$(VERSION).$(BUILD) forgewrapper.go

	sed -i 's/@NAME@/$(NAME)/g' $(APP_DIR)/Contents/Info.plist	
	sed -i 's/@VERSION@/$(VERSION)/g' $(APP_DIR)/Contents/Info.plist
	sed -i 's/@EXE@/$(NAME)_OSX-$(VERSION).$(BUILD)/g' $(APP_DIR)/Contents/Info.plist

	mv $(NAME)_OSX-$(VERSION).$(BUILD) $(APP_DIR)/Contents/MacOs/$(NAME)_OSX-$(VERSION).$(BUILD)

	tar -zcvf $(BIN_DIR)/FPL_OSX-$(VERSION).$(BUILD).tar.gz $(APP_DIR)

linux:
	GOOS=linux GOARCH=386 go build -v -o $(NAME)_Linux-$(VERSION).$(BUILD)
	
	mv $(NAME)_Linux-$(VERSION).$(BUILD) $(BIN_DIR)/$(NAME)_Linux-$(VERSION).$(BUILD)

	tar -zcvf $(BIN_DIR)/FPL_Linux-$(VERSION).$(BUILD).tar.gz  $(BIN_DIR)/$(NAME)_Linux-$(VERSION).$(BUILD)

all: bindir windows osx linux
