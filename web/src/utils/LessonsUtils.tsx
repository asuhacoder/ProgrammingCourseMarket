import axios from 'axios';
import { User, Lesson } from '../config/Type';

const compare = (a: Lesson, b: Lesson): number => {
    if (a.sequence_number < b.sequence_number) {
        return -1;
    } else if (a.sequence_number > b.sequence_number) {
        return 1;
    } else {
        return 0;
    }
}

export const SortLessons = (lessons: Lesson[]): Lesson[] => {
    lessons.sort(compare);
    return lessons;
}

export const NumberingLessons = (lessons: Lesson[]): Lesson[] => {
    let tmp = []
    for (let i = 0; i < lessons.length; i++) {
      let lesson:Lesson = Object.assign({}, lessons[i]);
      lesson.sequence_number = i;
      tmp.push(lesson)
    }
    return tmp
}

export const UpdateLessons = (courseID: string, user: User, lesson: Lesson): void => {
    axios.put(`http://localhost:8080/api/v1/lessons/${lesson.uuid}`, {
        token: window.localStorage.getItem('programming-course-market'),
        user_id: user.uuid,
        course_id: courseID,
        sequence_number: lesson.sequence_number,
        title: lesson.title,
        introduction: lesson.introduction,
        body: lesson.body,
        default_code: lesson.default_code,
        answer_code: lesson.answer_code,
        language: lesson.language,
      })
        .then((response) => {
          console.log(response);
        }, (error) => {
          console.log(error);
        });
}