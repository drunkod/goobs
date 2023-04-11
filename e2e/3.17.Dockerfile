FROM alpine:3.17 AS build
ENV obs_studio_version    28.1.2
# ENV obs_websocket_version 4.9.1

RUN apk add -U git build-base cmake
RUN apk add -U qt5-qtbase-dev ffmpeg-dev curl-dev wayland-dev x264-dev qt5-qtsvg-dev
RUN apk add -U  jansson-dev luajit-dev swig python3-dev libx11-dev

# install obs (https://obsproject.com/wiki/install-instructions#linux-portable-mode-all-distros)
WORKDIR /work
RUN git clone --recursive https://github.com/obsproject/obs-studio.git
WORKDIR /work/obs-studio/build
RUN git checkout "$obs_studio_version"
RUN cmake -DUNIX_STRUCTURE=0 -DCMAKE_INSTALL_PREFIX=/opt/obs -DBUILD_BROWSER=OFF -DBUILD_VST=OFF DENABLE_PIPEWIRE=OFF ..
RUN make -j4
RUN make install

FROM alpine:3.17
COPY --from=build /opt /opt
COPY --from=build /lib/libmount.so.1 /lib/libmount.so.1
COPY --from=build /lib/libblkid.so.1 /lib/libblkid.so.1
COPY --from=build /usr/lib /usr/lib
COPY basic.ini /root/.config/obs-studio/basic/profiles/Untitled/basic.ini
WORKDIR /opt/obs/bin/64bit/

# --init doesn't seem to work within a GH actions service container, so we add
# the init system ourselves to essentially always run with --init
ENV tini_version v0.19.0
ADD https://github.com/krallin/tini/releases/download/${tini_version}/tini-static /tini
RUN chmod +x /tini
ENTRYPOINT ["/tini", "--"]

RUN apk add -U xvfb-run mesa-dri-gallium

CMD ["xvfb-run", "-s", "-screen 0 800x600x24", "./obs"]
