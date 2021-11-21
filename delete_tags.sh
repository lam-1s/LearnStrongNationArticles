#!/bin/bash
set -e
git fetch --tags
TAGS=( $(git tag -l "*" | sort -r) )

KEEP_LATEST=50

LATEST_TAGS=()


for TAG in ${TAGS[@]}; do
    if [[ ${#LATEST_TAGS[@]} -lt ${KEEP_LATEST} ]]; then
        LATEST_TAGS+=( "$TAG" )
    fi
done

for TAG in ${LATEST_TAGS[@]}; do
    TAGS=( "${TAGS[@]/$TAG}" )
done

for TAG in ${TAGS[@]}; do
    echo "Deleting ${TAG}"
    hub release delete "${TAG}"
    git tag -d "${TAG}"
done

git push --tags --prune
