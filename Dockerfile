FROM zikalino/autorest.cli.base
MAINTAINER zikalino

#ADD . /autorest.cli
# RUN cd /autorest.cli; npm install
ADD autorest.cli /autorest.cli
RUN cd /autorest.cli; npm install
ADD magic-modules /magic-modules
RUN bash -l -c 'cd /magic-modules;gem install bundler:2.0.1; bundle update --bundler; bundle install'
ADD terraform/run.sh /run.sh
RUN chmod +x /run.sh
ENTRYPOINT ["bash", "/run.sh"]
