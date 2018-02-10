'use strict';

const Generator = require('yeoman-generator');

module.exports = class extends Generator {
  prompting() {
    const prompts = [{
      type: 'input',
      name: 'package',
      message: 'package name (ex. github.com/suzuki-shunsuke/sample-app)'
    }];

    return this.prompt(prompts).then(answers => {
      this.answers = answers;
    });
  }

  writing() {
    ['main.go', 'controllers/healthCheck.go',
     'constants/logTypes/logTypes.go'].forEach(key => {
      this.fs.copyTpl(
        this.templatePath(key),
        this.destinationPath(key),
        {package: this.answers.package});
    });
  }
};
