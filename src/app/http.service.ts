import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  private url = 'http://localhost:9000/walmart'

  constructor(private http: HttpClient) { }

  getWalmartItems(query : String) {
    return this.http.post(this.url, query);
  }
}
