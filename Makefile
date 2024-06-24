default: all

.PHONY: all
all:
	make -C channel
	make -C cli
	make -C error-handling-article
	make -C ginkgo-hello
	make -C greet
	make -C minimax
	make -C romannumeral
	make -C signal
	make -C stringcalc
	make -C time

#TODO KDK: Add test and make them all pass

.PHONY: tidy
tidy:
	make -C channel tidy
	make -C cli tidy
	make -C error-handling-article tidy
	make -C ginkgo-hello tidy
	make -C greet tidy
	make -C minimax tidy
	make -C romannumeral tidy
	make -C signal tidy
	make -C stringcalc tidy
	make -C time tidy
