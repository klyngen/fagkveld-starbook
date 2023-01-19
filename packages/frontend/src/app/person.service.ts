import { Injectable } from '@angular/core';
import { combineLatest, map, merge, Observable, Subject, switchMap, tap } from 'rxjs';
import { HttpService } from './http.service';
import { Person } from './models/person';
import { StarService } from './star.service';

@Injectable({
  providedIn: 'root'
})
export class PersonService {

  persons$: Observable<Person[]>;

  private personTrigger = new Subject();

  constructor(private httpClient: HttpService, readonly starService: StarService) {
    //this.persons$ = this.personTrigger.pipe(startWith(this.httpClient.fetchPersons()))
    this.persons$ = combineLatest([this.personObservable(), starService.starObservable()]).pipe(map(([persons, stars]) => {
        persons.forEach(person => {
          person.starCount = stars.filter(star => star.PersonID === person.ID).length;
        });
      console.log(persons);

        return persons;
      })
    );
  }

  createPerson(name: string, picture: string): Observable<void> {
    return this.httpClient.postPerson(name, picture).pipe(tap(() => this.personTrigger.next(null)));
  }

  private personObservable(): Observable<Person[]> {
    return merge(
      this.httpClient.fetchPersons(),
      this.personTrigger.pipe(switchMap(() => this.httpClient.fetchPersons()))
    );
  }


}
