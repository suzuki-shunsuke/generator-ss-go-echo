'use strict';

const Generator = require('yeoman-generator');

module.exports = class extends Generator {
  prompting() {
    const prompts = [{
      type: 'input',
      name: 'package',
      message: 'package name'
    }];

    return this.prompt(prompts).then(answers => {
      this.answers = answers;
    });
  }

  writing() {
    this.fs.copyTpl(
      this.templatePath('main.go'),
      this.destinationPath('main.go'),
      {package: this.answers.package});
    ['configs', 'controllers', 'models'].forEach(key => {
      this.fs.copy(
        this.templatePath(key),
        this.destinationPath(key));
    });
  }
};
