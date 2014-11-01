serve:
	bundle exec jekyll serve -w

dot:
	./dot.sh

update:
	bundle update

startingfromscratch:
	#https://help.github.com/articles/using-jekyll-with-pages
	#install ruby 1.9.3 or 2.0.0
	sudo gem install bundler
	sudo bundle install
