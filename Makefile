.DEFAULT_GOAL=all

BIN = hello_assert_bot
OBJ = \
	hello_assert_bot.cpp.o \
	IRCClient.cpp.o \
	IRCHandler.cpp.o \
	IRCSocket.cpp.o

CFLAGS = \
	-std=c++17 \
	-Wall \
	-Wextra \
	-Wno-unused-parameter \
	-Wno-missing-field-initializers \
	-Wno-sign-compare \
	-Wno-array-bounds \
	-g

LDFLAGS =

define HELP
build	Build program.
clean	Remove build files.
endef

export HELP
all:
	@echo "$$HELP"

build: $(BIN)

clean:
	rm -v $(BIN) $(OBJ)

$(BIN): $(OBJ)
	$(CXX) -o $@ $(OBJ) $(LDFLAGS)

%.cpp.o: %.cpp
	$(CXX) -c $(CFLAGS) $< -o $@
