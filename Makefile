.PHONY: default
default: all
	@:

## Configuration

sub_projects := \
	channel \
	cli \
	error-handling-article \
	ginkgo-hello \
	greet \
	minimax \
	signal \
	time \

.PHONY: debug
debug:
	$(info sub_projects: $(sub_projects))
	@:

## Standard targets

.PHONY: all
all:
	$(foreach sub_project,$(sub_projects),make -C $(sub_project) all;)

.PHONY: clean
clean:
	$(foreach sub_project,$(sub_projects),make -C $(sub_project) clean;)

.PHONY: test
test:
	$(foreach sub_project,$(sub_projects),make -C $(sub_project) test;)

## Top-level targets

.PHONY: tidy
tidy:
	$(foreach sub_project,$(sub_projects),make -C $(sub_project) tidy;)
