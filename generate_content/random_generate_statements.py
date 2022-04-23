# /bin/python3
import faker

str = faker.Faker(locale="zh_CN").sentence(nb_words=1)
print(len(str))
print(str)
# print(faker.Faker(locale="zh_CN").text(max_nb_chars=1500))
