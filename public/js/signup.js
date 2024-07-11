document.getElementById("form").addEventListener('submit', async (event) => {
    event.preventDefault();

    const keys = ["name", "surname", "username", "password"];

    keys.forEach(key => {
        document.getElementById(key).style.borderColor = "#ced4da";
        document.getElementById(`${key}error`).innerText = "";
    });

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

    res = await res.json()

    if (res.success) {
        window.location.href = "/login"
    } else {
        Object.keys(res.errors).forEach(key => {
            document.getElementById(key).style.borderColor = "rgb(255, 0, 0)";
            document.getElementById(`${key}error`).innerText = res.errors[key];
        });
    };
});
