#!/bin/bash
cp -rf --verbose /source/dist /autorest.cli
cp -rf --verbose /source/input /autorest.cli
#cd /magic-modules; bundle update --bundler; bundle
for d in $1/*.md; do
  autorest --cli --use=/autorest.cli --mm --intermediate --python-sdks-folder=/generated --output-folder=/generated $d
done
src="/generated"
if [ ! -d "/generated/magic-modules-input" ];then
  src="/generated"
else
  src="/generated/magic-modules-input"
fi
echo $src
cd /magic-modules; 
for resource in $src/*; do
  echo $resource
  bundle exec compiler.rb -e terraform -c azure -p $resource -o /mmoutput
done
