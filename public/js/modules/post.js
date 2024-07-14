function accommodatePosts(tag, posts) {
    function like(likes) {
        if (likes === null) return "no likes";
        return `${likes.length} like${likes.length > 1 ? 's' : ''}`
    }

    function comment(comments) {
        if (comments === null) return "no comment";
        return `${comments.length} comment${comments.length > 1 ? 's' : ''}`
    }

    tag.innerHTML = ""
    posts.forEach(post => {
        tag.innerHTML += `
        <div class="card">
            <a href="/post/${post.id}">
                <div class="card-body">

                    <h5 class="card-title">${post.user.name} ${post.user.surname}</h5>

                    <h6 class="card-subtitle text-muted">@${post.user.username}</h6>

                    <p class="card-text">${post.text}</p>

                    <h6 class="card-subtitle text-muted">${like(post.likes)} ${comment(post.comments)}</h6>
                    
                </div>
            </a>
        </div>  
    `;
    });
    // <h6 class="card-subtitle text-muted">${like(post.likesAmount)} ${comment(post.commentsAmount)}</h6>
} 