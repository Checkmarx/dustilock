package dependencies

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

type t struct {
	content  string
	expected []string
}

var npmTests = []t{
	{
		content:  `{"dependencies":{"angular":"123"}}`,
		expected: []string{"angular"},
	},
	{
		content:  `{"dependencies":{"@angular/core":"123"}}`,
		expected: []string{},
	},
}
var pythonTests = []t{

	{
		content:  `ipdb`,
		expected: []string{"ipdb"},
	},
	{
		content: `a_b>=1.0
	`,
		expected: []string{"a_b"},
	},
	{
		content: `
	
	
	`,
		expected: []string{},
	},
	{
		content: `-e bzr+http://bzr.myproject.org/MyProject/trunk#egg=MyProject
	-e bzr+sftp://user@myproject.org/MyProject/trunk#egg=MyProject
	-e bzr+ssh://user@myproject.org/MyProject/trunk#egg=MyProject
	-e bzr+ftp://user@myproject.org/MyProject/trunk#egg=MyProject
	-e bzr+https://bzr.myproject.org/MyProject/trunk@2019#egg=MyProject
	-e bzr+http://bzr.myproject.org/MyProject/trunk@v1.0#egg=MyProject
	
	bzr+http://bzr.myproject.org/MyProject/trunk#egg=MyProject
	bzr+sftp://user@myproject.org/MyProject/trunk#egg=MyProject
	bzr+ssh://user@myproject.org/MyProject/trunk#egg=MyProject
	bzr+ftp://user@myproject.org/MyProject/trunk#egg=MyProject
	bzr+https://bzr.myproject.org/MyProject/trunk@2019#egg=MyProject
	bzr+http://bzr.myproject.org/MyProject/trunk@v1.0#egg=MyProject
	`,
		expected: []string{"myproject"},
	},
	{
		content: `-e git+git://git.myproject.org/MyProject#egg=MyProject
	-e git+https://git.myproject.org/MyProject#egg=MyProject
	-e git+ssh://git.myproject.org/MyProject#egg=MyProject
	-e git+ssh://git@git.myproject.org/MyProject#egg=MyProject
	-e git://git.myproject.org/MyProject.git@master#egg=MyProject
	-e git://git.myproject.org/MyProject.git@v1.0#egg=MyProject
	-e git://git.myproject.org/MyProject.git@da39a3ee5e6b4b0d3255bfef95601890afd80709#egg=MyProject
	
	git+git://git.myproject.org/MyProject#egg=MyProject
	git+https://git.myproject.org/MyProject#egg=MyProject
	git+ssh://git.myproject.org/MyProject#egg=MyProject
	git://git.myproject.org/MyProject.git@master#egg=MyProject
	git://git.myproject.org/MyProject.git@v1.0#egg=MyProject
	git://git.myproject.org/MyProject.git@da39a3ee5e6b4b0d3255bfef95601890afd80709#egg=MyProject
	`,
		expected: []string{"myproject"},
	},
	{
		content: `# Copyright (c) 2012, Crate and individual contributors.
	# All rights reserved.
	# 
	# Redistribution and use in source and binary forms, with or without
	# modification, are permitted provided that the following conditions are met:
	# 
	# 1. Redistributions of source code must retain the above copyright notice, this
	#    list of conditions and the following disclaimer.
	# 
	# 2. Redistributions in binary form must reproduce the above copyright notice,
	#    this list of conditions and the following disclaimer in the documentation
	#    and/or other materials provided with the distribution.
	# 
	# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
	# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
	# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
	# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
	# ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
	# (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
	# LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
	# ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
	# (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
	# SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
	
	
	--extra-index-url=http://dist.pinaxproject.com/dev/
	
	Babel==0.9.6
	Django==1.4
	Jinja2==2.6
	MarkupSafe==0.15
	South==0.7.5
	amqplib==1.0.2
	anyjson==0.3.3
	bleach==1.1.1
	boto==2.3.0
	celery==2.5.5
	celery-haystack==0.4
	certifi==0.0.8
	chardet==1.0.1
	-e git+https://github.com/crateio/crate.web.git#egg=crate.web
	-e git+https://github.com/crateio/crate.pypi.git#egg=crate.pypi
	eventlet==0.9.16
	django-admin-tools==0.4.1
	django-appconf==0.5
	django-celery==2.5.5
	-e git+https://github.com/toastdriven/django-haystack.git#egg=django-haystack
	django-hosts==0.4.2
	django-jsonfield==0.8.7
	django-model-utils==1.1.0
	django-picklefield==0.2.1
	django-redis-cache==0.9.3
	django-secure==0.1.2
	django-social-auth==0.6.9
	django-staticfiles==1.2.1
	django-storages==1.1.4
	django-tastypie==0.9.11
	-e git://github.com/dstufft/django-user-accounts.git#egg=django-user-accounts
	django-uuidfield==0.2
	docutils==0.9.1
	greenlet==0.4.0
	gunicorn==0.14.5
	html5lib==0.95
	httplib2==0.7.4
	isoweek==1.2.0
	jingo==0.4
	-e git+https://github.com/dstufft/jutils.git#egg=jutils
	kombu==2.1.8
	lxml==2.3.4
	mimeparse==0.1.3
	newrelic==1.2.1.265
	oauth2==1.5.211
	oauthlib==0.1.3
	pinax-utils==1.0b1.dev3
	psycopg2==2.4.5
	py-bcrypt==0.2
	pyasn1==0.1.3
	-e git+https://github.com/toastdriven/pyelasticsearch.git#egg=pyelasticsearch
	python-dateutil==1.5
	python-openid==2.2.5
	pytz==2012c
	PyYAML==3.10
	raven==1.7.6
	redis==2.4.12
	requests==0.12.1
	rsa==3.0.1
	-e git+https://github.com/toastdriven/saved_searches.git#egg=saved_searches
	simplejson==2.5.2
	slumber==0.4.2
	uuid==1.30
	`,
		expected: []string{"jinja2", "eventlet", "south", "crate.pypi", "markupsafe", "django-secure", "lxml", "certifi", "celery", "django-appconf", "pinax-utils", "celery-haystack", "python-dateutil", "anyjson", "docutils", "django-model-utils", "rsa", "mimeparse", "gunicorn", "html5lib", "django-haystack", "django-redis-cache", "redis", "saved_searches", "oauthlib", "uuid", "pyelasticsearch", "chardet", "python-openid", "django-staticfiles", "greenlet", "jutils", "babel", "raven", "pyyaml", "simplejson", "slumber", "django-uuidfield", "isoweek", "pyasn1", "crate.web", "requests", "kombu", "django-storages", "psycopg2", "bleach", "django-picklefield", "httplib2", "django-admin-tools", "boto", "django-celery", "oauth2", "pytz", "django-hosts", "django", "newrelic", "jingo", "django-tastypie", "amqplib", "django-user-accounts", "django-jsonfield", "py-bcrypt", "django-social-auth"},
	},
	{
		content: `-e svn+svn://svn.myproject.org/svn/MyProject#egg=MyProject
	--editable git://git.myproject.org/MyProject.git@da39a3ee5e6b4b0d3255bfef95601890afd80709#egg=MyProject
	--editable hg+http://hg.myproject.org/MyProject/@special_feature#egg=MyProject
	-e bzr+lp://MyProject#egg=MyProject
	`,
		expected: []string{"myproject"},
	},
	{
		content: `-e svn+svn://svn.myproject.org/svn/MyProject#egg=MyProject
	-e svn+http://svn.myproject.org/svn/MyProject/trunk@2019#egg=MyProject
	
	svn+svn://svn.myproject.org/svn/MyProject#egg=MyProject
	svn+http://svn.myproject.org/svn/MyProject/trunk@2019#egg=MyProject
	`,
		expected: []string{"myproject"},
	},
	{
		content: `-e git+https://github.com/davidfischer/requirements-parser.git#egg=requirements
	Django >=1.5, <1.6
	numpy
	DocParser [PDF]
	`,
		expected: []string{"requirements", "django", "docparser", "numpy"},
	},
	{
		content: `-e svn+svn://svn.myproject.org/svn/MyProject#egg=MyProject # Test comment
	--editable git://git.myproject.org/MyProject.git@da39a3ee5e6b4b0d3255bfef95601890afd80709#egg=MyProject # Test comment
	--editable hg+http://hg.myproject.org/MyProject/@special_feature#egg=MyProject # Test comment
	-e bzr+lp://MyProject#egg=MyProject # Test comment
	`,
		expected: []string{"myproject"},
	},
	{
		content: `git+https://git.myproject.org/MyProject.git@v0.1#egg=MyProject[security]
	`,
		expected: []string{"myproject"},
	},
	{
		content: `file:///path/to/your/lib/project#egg=MyProject
	file://../../lib/project#egg=MyProject
	path/to/SomeProject#egg=SomeOtherProject
	-e path/to/AnotherProject#egg=AnotherProject
	`,
		expected: []string{"anotherproject", "myproject", "someotherproject"},
	},
	{
		content: `-r recursive_3.txt
	Jinja`,
		expected: []string{"jinja"},
	},
	{
		content: `# This is a comment
	`,
		expected: []string{},
	},
	{
		content: `# Copyright (c) 2011 Charles Leifer, Eric Holscher, Bobby Grace
	# 
	# Permission is hereby granted, free of charge, to any person
	# obtaining a copy of this software and associated documentation
	# files (the "Software"), to deal in the Software without
	# restriction, including without limitation the rights to use,
	# copy, modify, merge, publish, distribute, sublicense, and/or sell
	# copies of the Software, and to permit persons to whom the
	# Software is furnished to do so, subject to the following
	# conditions:
	# 
	# The above copyright notice and this permission notice shall be
	# included in all copies or substantial portions of the Software.
	# 
	# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
	# OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	# HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
	# WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
	# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
	# OTHER DEALINGS IN THE SOFTWARE.
	
	
	
	# Pypi ftw.
	Distutils2==1.0a3
	Sphinx==1.1.2
	Unipath==0.2.1
	bzr==2.5b4
	celery-haystack==0.6.2
	celery==3.0.9
	django-celery==3.0.9
	django-extensions==0.7.1
	django-guardian==1.0.4
	django-kombu==0.9.4
	django-profiles==0.2
	django-secure==0.1.2
	django==1.4.2
	docutils==0.8.1
	github2==0.5.2
	httplib2==0.7.2
	mercurial==2.4
	mimeparse==0.1.3
	redis==2.7.1
	simplejson==2.3.0
	slumber==0.4.2
	south==0.7.6
	sphinx-http-domain==0.2
	unittest-xml-reporting==1.3.1
	
	
	# Pegged git requirements
	git+git://github.com/toastdriven/django-haystack@259274e4127f723d76b893c87a82777f9490b960#egg=django_haystack
	git+git://github.com/alex/django-filter.git#egg=django-filter
	git+git://github.com/alex/django-taggit.git@36f6dabcf10e27c7d9442a94243d4189f2a4f121#egg=django_taggit-dev
	git+git://github.com/ericflo/django-pagination.git@e5f669036c#egg=django_pagination-dev
	git+git://github.com/nathanborror/django-basic-apps.git@171fdbe21a0dbbb38919a383cc265cb3cbc73771#egg=django_basic_apps-dev
	git+git://github.com/nathanborror/django-registration.git@dc0b564b7bfb79f58592fe8ad836729a85ec17ae#egg=django_registration-dev
	git+git://github.com/toastdriven/django-tastypie.git@c5451b90b18b0cb64841b2276d543230d5f58231#egg=django_tastypie-dev
	
	`,
		expected: []string{"south", "django-secure", "mercurial", "django-filter", "django_pagination-dev", "celery", "django-profiles", "celery-haystack", "docutils", "mimeparse", "redis", "sphinx-http-domain", "django-extensions", "django_tastypie-dev", "simplejson", "slumber", "sphinx", "django_registration-dev", "unipath", "github2", "django-guardian", "httplib2", "django-celery", "django_haystack", "django", "unittest-xml-reporting", "django_basic_apps-dev", "django-kombu", "django_taggit-dev", "bzr", "distutils2"},
	},
	{
		content: `PickyThing<1.6,>1.9,!=1.9.6,<2.0a0,==2.4c1
	`,
		expected: []string{"pickything"},
	},
	{
		content: `MyPackage[PDF]==3.0
	Fizzy [foo, bar]
	`,
		expected: []string{"mypackage", "fizzy"},
	},
	{
		content: `z3c.checkversions==0.4.1
	MyPackage
	Framework==0.9.4
	Library>=0.2
	`,
		expected: []string{"mypackage", "framework", "z3c.checkversions", "library"},
	},
	{
		content: `# Copyright (c) 2011 Charles Leifer, Eric Holscher, Bobby Grace
	# 
	# Permission is hereby granted, free of charge, to any person
	# obtaining a copy of this software and associated documentation
	# files (the "Software"), to deal in the Software without
	# restriction, including without limitation the rights to use,
	# copy, modify, merge, publish, distribute, sublicense, and/or sell
	# copies of the Software, and to permit persons to whom the
	# Software is furnished to do so, subject to the following
	# conditions:
	# 
	# The above copyright notice and this permission notice shall be
	# included in all copies or substantial portions of the Software.
	# 
	# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
	# OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
	# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	# HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
	# WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
	# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
	# OTHER DEALINGS IN THE SOFTWARE.
	
	
	psycopg2
	gunicorn
	pysolr
	python-memcached
	dnspython
	`,
		expected: []string{"dnspython", "gunicorn", "pysolr", "python-memcached", "psycopg2"},
	},
	{
		content: `req==1.0    # comment
	`,
		expected: []string{"req"},
	},
	{
		content: `-e git+git://git.myproject.org/MyProject#egg=MyProject&subdirectory=setup
	git+git://git.myproject.org/MyProject#egg=MyProject&subdirectory=setup
	`,
		expected: []string{"myproject"},
	},
	{
		content: `http://pypi.python.org/packages/source/p/pytz/pytz-2016.4.tar.gz#md5=a3316cf3842ed0375ba5931914239d97
	http://pypi.python.org/packages/source/f/flask/Flask-0.11.1.tar.gz#sha256=1212aaf123911123babc024abaa&egg=Flask
	`,
		expected: []string{"flask"},
	},
	{
		content: `-e hg+http://hg.myproject.org/MyProject#egg=MyProject
	-e hg+https://hg.myproject.org/MyProject#egg=MyProject
	-e hg+ssh://hg.myproject.org/MyProject#egg=MyProject
	-e hg+http://hg.myproject.org/MyProject@da39a3ee5e6b#egg=MyProject
	-e hg+http://hg.myproject.org/MyProject@2019#egg=MyProject
	-e hg+http://hg.myproject.org/MyProject@v1.0#egg=MyProject
	-e hg+http://hg.myproject.org/MyProject@special_feature#egg=MyProject
	
	hg+http://hg.myproject.org/MyProject#egg=MyProject
	hg+https://hg.myproject.org/MyProject#egg=MyProject
	hg+ssh://hg.myproject.org/MyProject#egg=MyProject
	hg+http://hg.myproject.org/MyProject@da39a3ee5e6b#egg=MyProject
	hg+http://hg.myproject.org/MyProject@2019#egg=MyProject
	hg+http://hg.myproject.org/MyProject@v1.0#egg=MyProject
	hg+http://hg.myproject.org/MyProject@special_feature#egg=MyProject
	`,
		expected: []string{"myproject"},
	},
}

func TestParsePythonRequirements(t *testing.T) {
	for _, item := range pythonTests {
		reader := bufio.NewReader(strings.NewReader(item.content))
		parsedPackages := ParsePythonRequirements(reader)

		sort.Strings(item.expected)
		sort.Strings(parsedPackages)

		expectedPackagesString := strings.Join(item.expected, ",")
		parsedPackagesString := strings.Join(parsedPackages, ",")

		if expectedPackagesString != parsedPackagesString {
			t.Errorf(`"%v" != "%v"`, expectedPackagesString, parsedPackagesString)
		}
	}
}

func TestParseNpmRequirements(t *testing.T) {
	for _, item := range npmTests {
		reader := bufio.NewReader(strings.NewReader(item.content))
		parsedPackages, _ := ParsePackagesJsonFile(reader)

		sort.Strings(item.expected)
		sort.Strings(parsedPackages)

		expectedPackagesString := strings.Join(item.expected, ",")
		parsedPackagesString := strings.Join(parsedPackages, ",")

		if expectedPackagesString != parsedPackagesString {
			t.Errorf(`"%v" != "%v"`, expectedPackagesString, parsedPackagesString)
		}
	}
}
