from setuptools import setup, find_packages

setup(
    name='your-project-name',
    version='1.0.0',
    author='Your Name',
    author_email='your-email@example.com',
    description='Description of your project',
    packages=find_packages(),
    install_requires=[
        'attrs==23.1.0',
        'blinker==1.6.2',
        'certifi==2023.5.7',
        'click==8.1.4',
        'exceptiongroup==1.1.2',
        'Flask==2.3.2',
        'h11==0.14.0',
        'idna==3.4',
        'importlib-metadata==6.8.0',
        'itsdangerous==2.1.2',
        'Jinja2==3.1.2',
        'jsonify==0.5',
        'MarkupSafe==2.1.3',
        'outcome==1.2.0',
        'PySocks==1.7.1',
        'selenium==4.10.0',
        'sniffio==1.3.0',
        'sortedcontainers==2.4.0',
        'trio==0.22.1',
        'trio-websocket==0.10.3',
        'urllib3==2.0.3',
        'Werkzeug==2.3.6',
        'wsproto==1.2.0',
        'zipp==3.15.0',
    ],
)
