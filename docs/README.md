# run natively using bundle
```
bundle install
bundle exec middleman server
```

## rbenv
You can install `rbenv` to manage your Ruby environments (just like pyenv for Python).

https://github.com/rbenv/rbenv

On OSX
```
brew install rbenv
echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> ~/.bash_profile
source ~/.bashrc
rbenv install 2.6.3
rbenv local 2.6.3
rbenv rehash
gem install bundler
```

# run with vagrant
```
vagrant up
```

You can now see the docs at http://localhost:4567
