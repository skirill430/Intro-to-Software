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
  error:any;

  constructor(private router:Router, private httpService: HttpService){}


  sendSignup(pageName:string) {
    console.log(this.LoginForm.controls.usernameControl.value);
    console.log(this.LoginForm.controls.passwordControl.value);
    
    this.httpService.sendSignupInfo(this.LoginForm.controls.usernameControl.value, this.LoginForm.controls.passwordControl.value).subscribe(response => {
      console.log(response);
      console.log(response.status);

      this.router.navigate([`${pageName}`]);
    }, error => {
      this.error = error;

      // check for user already exists
      if (error.status == 409)
        this.error = "User already exists";

      console.log(error);
    });
  }

  sendLogin(pageName:string) {
    console.log(this.LoginForm.controls.usernameControl.value);
    console.log(this.LoginForm.controls.passwordControl.value);

    // if (!this.error) {
    //   this.router.navigate([`${pageName}`]);
    // }

    this.httpService.sendLoginInfo(this.LoginForm.controls.usernameControl.value, this.LoginForm.controls.passwordControl.value).subscribe(response => {
      console.log(response);
      console.log(response.status);

      this.router.navigate([`${pageName}`]);
  }, error => {
      this.error = error;

      // check for invalid login
      if (error.status == 404)
        this.error = "User not found";

      console.log(error);
    });
  }

  // Routing function to take user to Home Page
  goToPage(pageName:string):void {
    if (!this.error) {
      this.router.navigate([`${pageName}`]);
    }
  }
}
