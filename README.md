Простой пример работы с K8S в кластере minikube.

Есть 2 небольших директории: 
    Первая - k8s-web-hello. 
    В данной директории создается простое приложение, которое просто по адресу: "/k8s" будет выводить на страницу привествтие.

    Вторая директория - k8s-web-to-nginx. 
    Здесь также создается простое приложение, которое по определенному адресу выводит приветствие, а также по адресу "/nginx", наше приложение будет обращаться внутри кластера K8S к сервису, у которого на подах будет запущено nginx приложение. И при обращении на этот адрес, будет просто выводиться приветственная страница nginx.

Подробности, о том как работают оба этих приложения, есть в их директориях в файле README.