import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpService } from '../http.service';
import { FormControl, FormGroup } from '@angular/forms';
import * as http from '../http.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  LoginForm: FormGroup = new FormGroup({
    usernameControl: new FormControl(''),
    passwordControl: new FormControl('')
  });
  title = 'Login Page';
  constructor(private router:Router, private httpService: HttpService){}


  sendSignup() {
    console.log(this.LoginForm.controls.usernameControl.value);
    console.log(this.LoginForm.controls.passwordControl.value);
    
    this.httpService.sendSignupInfo(this.LoginForm.controls.usernameControl.value, this.LoginForm.controls.passwordControl.value).subscribe(response => {
      console.log(response.status);
    })
  }

  sendLogin() {
    console.log(this.LoginForm.controls.usernameControl.value);
    console.log(this.LoginForm.controls.passwordControl.value);

    this.httpService.sendLoginInfo(this.LoginForm.controls.usernameControl.value, this.LoginForm.controls.passwordControl.value).subscribe(response => {
      console.log(response.status);
    })
  }

  // // function to take user to home page (not used)
  // goToPage(pageName:string):void {
  //   this.router.navigate([`${pageName}`]);
  // }
}
