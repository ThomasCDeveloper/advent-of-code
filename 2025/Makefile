.PHONY: run test clean init

run:
	@python3 $(DAY)/solution.py input

test:
	@python3 $(DAY)/solution.py test

clean:
	find . -name '*.pyc' -delete

init:
	@mkdir -p $(DAY)
	@touch $(DAY)/input.txt $(DAY)/test.txt
	@cp template/solution.py $(DAY)/solution.py
	@echo "done"
