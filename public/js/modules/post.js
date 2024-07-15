async function accommodatePosts(tag, posts) {
    function like(likes) {
        if (likes === null) return "no likes";
        return `${likes.length} like${likes.length > 1 ? 's' : ''}`
    };

    function comment(comments) {
        if (comments === null) return "no comment";
        return `${comments.length} comment${comments.length > 1 ? 's' : ''}`
    };

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
                `<button type="button" class="btn btn-secondary btn-sm" onclick="event.preventDefault(); unlike('${id}')">Unlike</button>` :
                `<button type="button" class="btn btn-primary btn-sm" onclick="event.preventDefault(); like('${id}')">Like</button>`;
        };
    };

    tag.innerHTML = ""
    for (post of posts) {
        console.log("fuckkkkkk")
        tag.innerHTML += `
        <div class="card">
            <a href="/post/${post.id}">
                <div class="card-body">

                    <h5 class="card-title">${post.user.name} ${post.user.surname}</h5>

                    <h6 class="card-subtitle text-muted"
                    onclick="event.preventDefault(); window.location.href = '/user/${post.user.username}'"
                    style="display: inline-block;">
                        @${post.user.username}
                    </h6>

                    <p class="card-text">${post.text}</p>

                    <h6 class="card-subtitle text-muted">${like(post.likes)} ${comment(post.comments)}</h6>

                    <div id="likebtns-${post.id}" style="display: inline-block;">
                        ${await likedCase(post.id)}
                    </div>
                    
                </div>
            </a>
        </div>  
        `;
    };
};

async function like(id) {
    let res = await fetch(`/api/post/${id}/like`, {
        method: 'PUT',
        headers: {
            "Content-Type": "application/json"
        },
    })

    res = await res.json();

    if (res.success) {
        document.getElementById(`likebtns-${id}`).innerHTML = `
            <button type="button" class="btn btn-secondary btn-sm" onclick="event.preventDefault(); unlike('${id}')">Unlike</button>
        `;
    } else {
        alert(res.message);
    }
};

async function unlike(id) {
    let res = await fetch(`/api/post/${id}/unlike`, {
        method: 'PUT',
        headers: {
            "Content-Type": "application/json"
        },
    })

    res = await res.json();

    if (res.success) {
        document.getElementById(`likebtns-${id}`).innerHTML = `
            <button type="button" class="btn btn-primary btn-sm" onclick="event.preventDefault(); like('${id}')">Like</button>
        `;
    } else {
        alert(res.message);
    }
};