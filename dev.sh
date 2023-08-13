#!/bin/bash

air &

air_pid=$!

cd front

bun run build2.js &

bun_pid=$!

trap onexit INT
function onexit() {
    kill $air_pid $bun_pid
}

bun run tailwind:dev
