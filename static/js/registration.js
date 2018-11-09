// Валидация имени
$('#first').blur(function() {
    var first = $('#first').val();
    if (first.length < 1 ) {
        $('#firstDesc').html("Введите имя");
        return false;
    }
});

$('#last').blur(function() {
    var last = $('#last').val();
    if (last.length < 1 ) {
        $('#lastDesc').html("Введите фамилию");
        return false;
    }
});

// Валидация почты
var email = $('#email');
var emailDesc = $("#emailDesc");
email.on('input', emailValid);
$('form').on('submit', emailValid);


function emailValid (e){
    var output = "ok";
    var err = false;
    var mail = $.trim(email.val());

    if(mail === '') {
        output = 'Введите почтовый адрес';
        err = true;
    } else {
        if( !validateEmail(mail) ) {
            output = 'Такой почты не может быть!';
            err = true;
        } else {
            $.ajax({
                url: "have",
                type: "post",
                dataType: "json",
                data: {email: mail},
                success: AjaxSucceeded,
                error: AjaxFailed
            });

            function AjaxSucceeded(data) {
                if (data == true) {
                    output = "Этот почтовый адрес уже используется";
                    emailDesc.text(output);
                    err = true

                }else if (data == false) {
                    output = "Этот почтовый адрес свободен";
                    emailDesc.text(output);
                }else{
                    alert("Chush!")
                }
            }

            function AjaxFailed(result) {
                output = "Ошибка сервера";
                err = true;
                alert(result.status + ' ' + result.statusText);
            }
        }
    }
    emailDesc.text(output);
    if(err) e.preventDefault();
}

function validateEmail($email) {
    var emailReg = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;
    return emailReg.test( $email );
}


// Валидация пароля
var passFields = $('.edinfopass'),
    validResult = $("#validpass");
passFields.on('input', comparingPasswords);
$('form').on('submit', comparingPasswords);

function comparingPasswords (e){
    var output = 'Поля паролей заполнены верно',
        err = false,
        p1 = $.trim(passFields.eq(0).val()),
        p2 = $.trim(passFields.eq(1).val());
    if(p1 == '' || p2 == '') {
        output = 'Заполните поля!';
        err = true;
    } else {
        if(p1.length < 4 || p2.length < 4) {
            output = 'Cлишком короткий пароль(минимум 4 знака)!';
            err = true;
        }else {
            if(p1 != p2) {
                output = 'Пароли не совпадают!';
                err = true;
            }
        }
    }

    validResult.text(output);
    if(err) e.preventDefault();
}

// Валидация возраста
var age = $('#age');
var ageDesk = $("#ageDesc");
age.on('input', ageValid);
$('form').on('submit', ageValid);

function ageValid(e) {
    var output = '',
        err = false,
        a = $.trim(age.val());
    if(a == '' ) {
        output = 'Укажите свой возраст!';
        err = true;
    } else {
        if(a < 0){
            output = 'Не может быть столько лет!';
            err = true;
        }else {
            if(a < 18) {
                output = 'Лицам моложе 18 лет запрещена регистрация на этом сайте!';
                err = true;
            }else {
                if ( a > 150 ) {
                    output = 'Вам нельзя сидеть за компьютером)'
                }
            }
        }
    }

    ageDesk.text(output);
    if(err) e.preventDefault();
}


function validateForm(form) {

    if (form.first.value === "") {
        $('#first').focus();
        alert("Введите имя");
        return false;
    }
    if (form.last.value === "") {
        $('#last').focus();
        alert("Введите фамилию");
        return false;
    }

    return true;
}