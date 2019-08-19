$(function () {

    loadUserList();

    function loadUserList() {
        let loginHeaders = new Headers();
        loginHeaders.append('Authorized', localStorage.getItem("token"));

        let request = new Request('/api/system/user-list', {
            method: 'GET',
            headers: loginHeaders
        });
        fetch(request)
            .then(response => response.json())
            .then(result => {
                if (result.code == 0) {
                    $("#table-body").html(GenerateHtml(result.data));
                } else {
                    alert(result.message);
                }
            })
    }

    function GenerateHtml(userList) {
        var html = '';
        for (let i = 0; i < userList.length; i++) {
            html += '<tr>\n' +
                '        <th scope="row">'+userList[i].ID+'</th>\n' +
                '        <td>'+userList[i].username+'</td>\n' +
                '        <td>'+userList[i].displayName+'</td>\n' +
                '        <td>'+userList[i].userStatus+'</td>\n' +
                '    </tr>';
        }
        return html;
    }
});