python3 -m venv env
source env/bin/activate
pip install django
python manage.py migrate
python manage.py runserver
