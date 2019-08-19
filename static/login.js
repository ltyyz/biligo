$(function () {
    $("#btn-login").click(function () {
        let username = $("#username").val();
        let password = $("#password").val();

        doLogin(username, password);

        return false;
    });

    function doLogin(username, password) {
        let loginHeaders = new Headers();
        loginHeaders.append('Content-Type', 'application/x-www-form-urlencoded');

        let request = new Request('/api/auth/login', {
            method: 'POST',
            mode: 'no-cors',
            body: "username=" + username + "&password=" + password,
            headers: loginHeaders
        });
        fetch(request)
            .then(response => response.json())
            .then(result => {
                if (result.code == 0) {
                    localStorage.setItem("token", result.data.token);
                    alert("登录成功!");
                    location.href = "userList.html";

                } else {
                    alert(result.message);
                }
            })
    }
});