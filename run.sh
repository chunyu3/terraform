cp -rf --verbose /source/dist /autorest.cli
cp -rf --verbose /source/input /autorest.cli
#cd /magic-modules; bundle update --bundler; bundle
for d in $1/*.md; do
  autorest --cli --use=/autorest.cli --mm --intermediate --python-sdks-folder=/generated --output-folder=/generated $d
  cd /magic-modules; bundle exec compiler.rb -e terraform -c azure -p /generated/alertsmanagementsmartdetectoralertrule -o /mmoutput
  break
done
