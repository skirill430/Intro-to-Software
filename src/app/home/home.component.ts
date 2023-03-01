import { Component } from '@angular/core';
import { AppComponent } from '../app.component'
import { HttpService } from '../http.service';
import * as http from '../http.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  title = 'Quick Shop';
  items : http.RootObject;
  search : string;
  showFiller = false;

  constructor(private httpService: HttpService) {}

  save() {
    console.log(this.search);
    this.httpService.getWalmartItems(this.search).subscribe(response => {
      console.log(response);
      this.items = response;
    })
  }
}