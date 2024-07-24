# Layers are ordered from the least likely to change to the most likely
# to change. By doing that, we reduce the number of layers to build and
# push to Artifactory as the non changing ones are reused.

#-----------------------------------------------------
# Use a golden image downloaded from the artifact repo
#-----------------------------------------------------
FROM @@IMAGE_ID@@

#-----------------------
# install basic packages
#-----------------------
RUN yum -y install iputils && yum clean all

# Use ENTRYPOINT to allow additional parameters to be passed on the command line:
# e.g. docker run uuid-gen -trace -debug
ENTRYPOINT ["/opt/uuid-gen/entrypoint.sh"]
CMD ["-trace"]

COPY entrypoint.sh /opt/uuid-gen/
#COPY assets/ /opt/kaloom/fps/assets
COPY uuid-gen-cli uuid-gen /opt/uuid-gen/bin/
