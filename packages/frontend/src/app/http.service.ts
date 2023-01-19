import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Person } from './models/person';
import { Star } from './models/star';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private httpClient: HttpClient) { }

  fetchPersons(): Observable<Person[]> {
    return this.httpClient.get<Person[]>("http://127.0.0.1:1337/person")
  }

  postPerson(name: string, image: string): Observable<void> {
    return this.httpClient.post<void>("http://127.0.0.1:1337/person", {
      name,
      image
    })
  }

  postStar(description: string, userId: number): Observable<void> {
    return this.httpClient.post<void>(`http://127.0.0.1:1337/star`, {
      description,
      personId: userId
    });
  }

  fetchStars(): Observable<Star[]> {
    return this.httpClient.get<Star[]>("http://127.0.0.1:1337/star");
  }
}
