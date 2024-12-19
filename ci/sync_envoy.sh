#!/bin/bash -e

set -o pipefail

MIRROR_MSG="Mirrored from xczhw/my-envoy"
SRCS=(envoy contrib)
GO_TARGETS=(@envoy_api//...)
IMPORT_BASE="github.com/xczhw/go-control-plane"
COMMITTER_NAME="xczhw"
COMMITTER_EMAIL="xczhw@outlook.com"
ENVOY_SRC_DIR="${ENVOY_SRC_DIR:-}"

if [[ -z "$ENVOY_SRC_DIR" ]]; then
    echo "DEBUG: ENVOY_SRC_DIR is not set, exiting..."
    exit 1
elif [[ ! -d "$ENVOY_SRC_DIR" ]]; then
    echo "DEBUG: ENVOY_SRC_DIR ($ENVOY_SRC_DIR) directory not found, exiting..."
    exit 1
else
    echo "DEBUG: ENVOY_SRC_DIR is set to $ENVOY_SRC_DIR"
fi

build_protos () {
    echo "DEBUG: Starting build_protos..."
    echo "DEBUG: Current directory: $(pwd)"
    echo "DEBUG: Changing directory to ${ENVOY_SRC_DIR}"
    cd "${ENVOY_SRC_DIR}" || exit 1
    echo "DEBUG: Running ./ci/do_ci.sh api.go"
    ./ci/do_ci.sh api.go
    cd - || exit 1
    echo "DEBUG: build_protos completed."
}

get_last_envoy_sha () {
    echo "DEBUG: Fetching last Envoy SHA..."
    git log --grep="$MIRROR_MSG" -n 1 | grep "$MIRROR_MSG" | tail -n 1 | sed -e "s#.*$MIRROR_MSG @ ##"
}

sync_protos () {
    echo "DEBUG: Starting sync_protos..."
    local src envoy_src
    for src in "${SRCS[@]}"; do
        envoy_src="${ENVOY_SRC_DIR}/build_go/${src}"
        echo "DEBUG: Syncing $src from $envoy_src"
        if [[ -d "$envoy_src" ]]; then
            rm -rf "$src"
            cp -a "$envoy_src" "$src"
            git add "$src"
        else
            echo "WARNING: $envoy_src not found, skipping..."
        fi
    done
    echo "DEBUG: Running make tidy-all"
    make tidy-all
    git add $(find -type f -name 'go.sum' -o -name 'go.mod')
    echo "DEBUG: sync_protos completed."
}

commit_changes () {
    echo "DEBUG: Starting commit_changes..."
    local last_envoy_sha changes changed

    # 检测变更文件
    changed="$(git diff HEAD --name-only | grep -v envoy/COMMIT || :)"
    if [[ -z "$changed" ]]; then
        echo "DEBUG: No changes detected, skipping commit."
        return
    fi

    # # 输出文件数量和示例文件
    echo "DEBUG: Total changed files: $(echo "$changed" | wc -l)"
    # echo "DEBUG: Sample changed files:"
    # echo "$changed" | head -n 10  # 只显示前 10 个文件

    # # 获取最新的 Envoy SHA
    # last_envoy_sha="$(get_last_envoy_sha)"
    # echo "DEBUG: Last Envoy SHA: ${last_envoy_sha}"

    # # 检测 Envoy 的变更
    # changes="$(git -C "${ENVOY_SRC_DIR}" rev-list "${last_envoy_sha}"..HEAD || :)"
    # echo "DEBUG: Detected changes from Envoy."

    # 获取最新的提交
    latest_commit="$(git -C "${ENVOY_SRC_DIR}" rev-list HEAD -n1)"
    echo "DEBUG: Latest commit: ${latest_commit}"
    echo "$latest_commit" > envoy/COMMIT

    # 配置 Git 用户
    git config user.email "$COMMITTER_EMAIL"
    git config user.name "$COMMITTER_NAME"
    echo "DEBUG: Git user.name is $(git config user.name)"
    echo "DEBUG: Git user.email is $(git config user.email)"

    # 添加更改到 Git
    echo "DEBUG: Staging files for commit..."
    git add envoy contrib || { echo "DEBUG: Failed to stage files"; exit 1; }

    # 提交更改
    echo "DEBUG: Committing changes..."
    git commit --allow-empty -s -m "${MIRROR_MSG} @ ${latest_commit}" || {
        echo "DEBUG: Git commit failed"
        exit 1
    }

    # 推送更改
    echo "DEBUG: Pushing changes to origin/main..."
    git remote -v
    git push origin master || {
        echo "DEBUG: Git push failed"
        exit 1
    }

    echo "DEBUG: commit_changes completed successfully."
}


build_protos
sync_protos
commit_changes
