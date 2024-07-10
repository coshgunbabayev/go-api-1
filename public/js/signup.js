document.getElementById("form").addEventListener('submit', async (event) => {
    event.preventDefault();
    
    const form = document.getElementById("form");
    const formData = new FormData(form);

    let res = await fetch("/api/user/signup", {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            name: formData.get("name"),
            surname: formData.get("surname"),
            username: formData.get("username"),
            password: formData.get("password"),
        })
    });

    console.log(await res.json());
});
