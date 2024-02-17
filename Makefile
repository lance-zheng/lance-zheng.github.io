sync: build lint
	@rsync -a --delete --exclude=".git" ./ ~/OneDrive/note
	@git config user.email 12312312343434@t12312312343434.com
	@git checkout --orphan empty_branch 
	@git add -A
	@git commit -am 'update' > /dev/null
	@git branch -D gh-pages
	@git branch -m gh-pages
	@git reflog expire --expire=90.days.ago --expire-unreachable=now --all
	@git push -f origin gh-pages
	@echo ""
	@echo "done."

wsl: build lint
	@git config user.email 12312312343434@t12312312343434.com
	@git checkout --orphan empty_branch 
	@git add -A
	@git commit -am 'update' > /dev/null
	@git branch -D gh-pages
	@git branch -m gh-pages
	@git reflog expire --expire=90.days.ago --expire-unreachable=now --all
	@git push -f origin gh-pages
	@echo ""
	@echo "done."


build: clean gobuild
	@./bin/app leetcode
	@./bin/app readme

gobuild:
	@cd generator && go mod tidy && go build -o ../bin/app ./

clean:
	@rm -rf ./bin
	@rm -rf ./sources/generated-sources 
	@mkdir ./sources/generated-sources 

lint:
	@command -v markdownlint > /dev/null || npm install -g markdownlint-cli
	@markdownlint '**/*.md' --disable MD013 MD033 MD045


newcode: gobuild
	@./bin/app newcode $(q)
	code ./sources/leetcode/$(q).md

cpimg: gobuild
	@./bin/app pasteimg $(q)

pull:
	git fetch --all && git reset --hard origin/gh-pages
