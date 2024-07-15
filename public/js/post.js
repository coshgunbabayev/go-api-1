const id = document.getElementById("id").innerText;
const postDetails = document.getElementById("postDetails");
const form = document.getElementById("form");
const commentsHead = document.getElementById("commentshead");
const comments = document.getElementById("comments");

async function getPostDetails() {
    let res = await fetch(`/api/post/${id}`, {
        method: 'GET',
        headers: {
            "Content-Type": "application/json"
        },
    })

    res = await res.json();

    if (res.success) {
        const post = res.post;

        function like(likes) {
            if (likes === null) return "no likes";
            return `${likes.length} like${likes.length > 1 ? 's' : ''}`
        }

        function comment(comments) {
            if (comments === null) return "no comment";
            return `${comments.length} comment${comments.length > 1 ? 's' : ''}`
        }

        async function likedCase(id) {
            let res = await fetch(`/api/post/${id}/like`, {
                method: 'GET',
                headers: {
                    "Content-Type": "application/json"
                },
            })

            res = await res.json();

            if (res.success) {
                return res.liked ?
                    `<button type="button" class="btn btn-secondary btn-sm" onclick="unlike('${id}')">Unlike</button>` :
                    `<button type="button" class="btn btn-primary btn-sm" onclick="like('${id}')">Like</button>`;
            }
        }

        postDetails.innerHTML = `
            <div class="card">
                <div class="card-body">

                    <h4 class="card-title">${post.user.name} ${post.user.surname}</h4>

                    <h6 class="card-subtitle text-muted"><a href="/user/${post.user.username}">@${post.user.username}</a></h6>

                    <h5 class="card-text">${post.text}</h5>

                    <h6 class="card-subtitle text-muted">${like(post.likes)} ${comment(post.comments)}</h6>

                    <div id="likebtns-${post.id}" style="display: inline-block;">
                        ${await likedCase(post.id)}
                    </div>

                    <div id="commentbtns" style="display: inline-block;">
                        <button type="button" class="btn btn-primary btn-sm" onclick="commentOpen()">Share comment</button>
                    </div>

                </div>
            </div>  
        `;

        commentsHead.innerHTML = `
            <div class="card">
                <div class="card-body">

                    <h5 class="card-title">
                        Comments
                    </h5>

                </div>
            </div>
        `;

        if (post.comments !== null) {
            accommodatePosts(comments, post.comments)
        } else {
            comments.innerHTML = `
                <div class="card">
                    <div class="card-body">

                        <p class="card-title">
                            Not found. Be the first to comment!
                        </p>

                    </div>
                </div>
            `;
        }
    } else if (res.message) {
        postDetails.innerHTML = `
            <div class="card">
                <div class="card-body">

                    <h5 class="card-title">
                        ${res.message}
                    </h5>

                </div>
            </div>
        `;
    };

}; getPostDetails();

function commentOpen() {
    document.getElementById("commentbtns").innerHTML = `
        <button type="button" class="btn btn-secondary btn-sm" onclick="commentClose()">Close</button>
    `;
    form.style.display = "flex";
}

function commentClose() {
    document.getElementById("commentbtns").innerHTML = `
        <button type="button" class="btn btn-primary btn-sm" onclick="commentOpen()">Share comment</button>
    `;
    form.style.display = "none";
}