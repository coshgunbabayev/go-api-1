const posts = document.getElementById("posts");

async function getUserDetails() {
    let res = await fetch("/api/post", {
        method: 'GET',
        headers: {
            "Content-Type": "application/json"
        },
    })

    res = await res.json();

    if (res.success) {
        res.posts.reverse();
        accommodatePosts(posts, res.posts)
    }
}; getUserDetails();