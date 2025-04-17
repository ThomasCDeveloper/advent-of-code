.PHONY: run test clean init

run:
	@python3 solutions/$(DAY)/solution.py input

test:
	@python3 solutions/$(DAY)/solution.py test

clean:
	find . -name '*.pyc' -delete

init:
	@mkdir -p solutions/$(DAY)
	@touch solutions/$(DAY)/input.txt solutions/$(DAY)/test.txt
	@cp template/solution.py solutions/$(DAY)/solution.py
	@echo "done"
