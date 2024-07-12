const username = document.getElementById("username").innerText;
const userDetails = document.getElementById("userdetails");

async function getUserDetails() {
    let res = await fetch(`/api/user/${username}`, {
        method: 'GET',
        headers: {
            "Content-Type": "application/json"
        },
    })

    res = await res.json();

    if (res.success) {
        const user = res.user;

        userDetails.innerHTML = `
            <div class="card">
                <div class="card-body">

                    <h4 class="card-title">
                        ${user.name} ${user.surname}
                    </h4>

                    <h6 class="card-subtitle text-muted">
                        @${user.username}
                    </h6>
                
                </div>
            </div>
        `;

    } else if (res.message) {
        userDetails.innerHTML = `
            <div class="card">
                <div class="card-body">

                    <h5 class="card-title">
                        ${res.message}
                    </h5>

                </div>
            </div>
        `;
    };

}; getUserDetails();