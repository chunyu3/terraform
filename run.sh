#!/bin/bash
echo $1
cp -rf --verbose /source/dist /autorest.cli
cp -rf --verbose /source/input /autorest.cli
#cd /magic-modules; bundle update --bundler; bundle
echo $1
whoami
if [ -z "$2" ]; then
  start=0
else
  start=$2
fi

if [ -z "$3" ]; then
  num=0
else
  num=$3
fi

cd /cli-input-tools
#for d in $1/*.md; do
 # autorest --cli --use=/autorest.cli --mm --intermediate --python-sdks-folder=/generated --output-folder=/generated $di
 # break
#done
RESULT_FOO=$(python3 list_readme_path_of_all_services.py $1)
#echo $RESULT_FOO
for i in ${RESULT_FOO}; do
    (( start-- ))
    if [ $start -gt 0 ]; then
        echo "skip"
        continue;
    fi
	echo $i;
    python3 ~/cli-input-tools/add_snippets_for_readme.py $i

	#s = $i;
	#echo $s
	#TAGS=`python -c 'import sys; import get_tags_from_readme; print get_tags_from_readme.get_tags_from_readme(sys.argv[1])', "$i"`
	TAGS=$(python3 get_tags_from_readme.py $i)
	echo $TAGS
	for t in ${TAGS}; do
  		autorest --cli --use=/autorest.cli --mm --intermediate --tag=$t --python-sdks-folder=/generated --output-folder=/generated $i
	done

    #trage the number of resource to generate
    (( num-- ))
    if [ $num -eq 0 ]; then
      break
    fi
done
src="/generated"
if [ ! -d "/generated/magic-modules-input" ];then
  src="/generated"
else
  src="/generated/magic-modules-input"
fi
echo $src
cd /magic-modules;
export PATH=/root/.rbenv/shims:$PATH
echo $PWD
echo $PATH
for resource in $src/*; do
  echo $resource
  f=$(basename -- $resource)
  echo $f
  for pkg in $resource/*; do
	  echo $pkg
	  pname=$(basename -- $pkg)
	  outputdir="/mmoutput/$f/$pname"
	  echo $outputdir
	  mkdir -p $outputdir
	  bundle exec compiler.rb -e terraform -c azure -p $pkg -o $outputdir
  done
done
