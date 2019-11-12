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
RESULT_FOO=$(python3 list_readme_path_of_updated_services.py $1 "remotes/origin/master" "remotes/terraform/master")
#echo $RESULT_FOO
echo "START Auto-cli Magic-Modules input files generation"
for i in ${RESULT_FOO}; do
    (( start-- ))
    if [ $start -gt 0 ]; then
        echo "skip"
        continue;
    fi
	echo "START $i";
    python3 /cli-input-tools/add_snippets_for_readme.py $i

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
echo "END Auto-cli Magic-Modules input files generation"

src="/generated"
if [ ! -d "/generated/magic-modules-input" ];then
  src="/generated"
else
  src="/generated/magic-modules-input"
fi
#echo $src
echo "START Magic-Module Generation"
cd /magic-modules;
export PATH=/root/.rbenv/shims:$PATH
echo $PWD
echo $PATH
for resource in $src/*; do
  echo "START $resource"
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
echo "END Magic-Module Generation"
