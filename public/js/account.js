function loginForm() {
    document.getElementById("form").innerHTML = `
        <h2 class="text-uppercase text-center mb-5" >Login</h2>
        <form id="loginform">

            <div class="form-outline mb-4">
                <label class="form-label" for="L_username">Your Username</label>
                <input type="text" id="L_username" name="username"
                    class="form-control form-control-lg" />

                <span id="L_usernameerror" class="error"></span>
            </div>

            <div class="form-outline mb-4">
                <label class="form-label" for="L_password">Password</label>
                <input type="password" id="L_password" name="password"
                    class="form-control form-control-lg" />

                <span id="L_passworderror" class="error"></span>
            </div>

            <div class="d-flex justify-content-center">
                <button type="submit" onclick="loginSbmt(event)"
                    class="btn btn-success btn-block btn-lg gradient-custom-4 text-body">
                    Login
                </button>
            </div>

            <p class="text-center text-muted mt-5 mb-0">Don't have an account?
                <button onclick="signupForm()" type="button" class="btn btn-primary btn-sm">Sign up here</button>
            </p>

        </form>
    `;
};


function signupForm() {
    document.getElementById("form").innerHTML = `
        <h2 class="text-uppercase text-center mb-5" >Create an account</h2>
        <form id="signupform">

            <div class="form-outline mb-4">
                <label class="form-label" for="S_name">Your Name</label>
                <input type="text" id="S_name" name="name" class="form-control form-control-lg" />

                <span id="S_nameerror" class="error"></span>
            </div>


            <div class="form-outline mb-4">
                <label class="form-label" for="S_surname">Your Surname</label>
                <input type="text" id="S_surname" name="surname"
                    class="form-control form-control-lg" />

                <span id="S_surnameerror" class="error"></span>
            </div>

            <div class="form-outline mb-4">
                <label class="form-label" for="S_username">Your Username</label>
                <input type="text" id="S_username" name="username"
                    class="form-control form-control-lg" />

                <span id="S_usernameerror" class="error"></span>
            </div>

            <div class="form-outline mb-4">
                <label class="form-label" for="S_password">Password</label>
                <input type="password" id="S_password" name="password"
                    class="form-control form-control-lg" />

                <span id="S_passworderror" class="error"></span>
            </div>

            <div class="d-flex justify-content-center">
                <button type="submit" onclick="signupSbmt(event)"
                    class="btn btn-success btn-block btn-lg gradient-custom-4 text-body">
                    Sign up
                </button>
            </div>

            <p class="text-center text-muted mt-5 mb-0">Have already an account?
                <button onclick="loginForm()" type="button" class="btn btn-primary btn-sm">Login here</button>
            </p>

        </form>
    `;
};

loginForm();

async function signupSbmt(event) {
    event.preventDefault();
    
    const keys = ["S_name", "S_surname", "S_username", "S_password"];

    keys.forEach(key => {
        document.getElementById(key).style.borderColor = "#ced4da";
        document.getElementById(`${key}error`).innerText = "";
    });

    const form = document.getElementById("signupform");
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
        loginForm()
    } else {
        Object.keys(res.errors).forEach(key => {
            document.getElementById(`S_${key}`).style.borderColor = "rgb(255, 0, 0)";
            document.getElementById(`S_${key}error`).innerText = res.errors[key];
        });
    };
};

async function loginSbmt(event) {
    event.preventDefault();
    
    const keys = ["L_username", "L_password"];

    keys.forEach(key => {
        document.getElementById(key).style.borderColor = "#ced4da";
        document.getElementById(`${key}error`).innerText = "";
    });

    const form = document.getElementById("loginform");
    const formData = new FormData(form);

    let res = await fetch("/api/user/login", {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: formData.get("username"),
            password: formData.get("password"),
        })
    });

    res =await res.json();

    if (res.success) {
        window.location.href = "/"
    } else {
        Object.keys(res.errors).forEach(key => {
            document.getElementById(`L_${key}`).style.borderColor = "rgb(255, 0, 0)";
            document.getElementById(`L_${key}error`).innerText = res.errors[key];
        });
    };
}