![dayp01](misc/images/_6ffb210b-602d-44f2-918d-1df3942f2222.jpeg)

# Contents

1 [Exercise 00: Reading](#ex00) \
2 [Exercise 01: Assessing Damage](#ex01)\
3 [Exercise 02: Dumps](#ex02)

<h2 id="ex00">Exercise 00: Reading</h2>

### XML:

```xml
<recipes>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>40 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>3</itemcount>
                <itemunit>cups</itemunit>
            </item>
            <item>
                <itemname>Vanilla extract</itemname>
                <itemcount>1.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>7</itemcount>
                <itemunit></itemunit> <!-- itemunit may be empty  -->
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit>pieces</itemunit>
            </item>
            <!-- Here can be more ingredients  -->
        </ingredients>
    </cake>
    <cake>
        <name>Blueberry Muffin Cake</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Baking powder</itemname>
                <itemcount>3</itemcount>
                <itemunit>teaspoons</itemunit>
            </item>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>0.5</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>cup</itemunit>
            </item>
            <!-- Here can be more ingredients  -->
        </ingredients>
    </cake>
    <!-- Here can be more cakes  -->
</recipes>
```
### JSON:

```json
{
  "cake": [
    {
      "name": "Red Velvet Strawberry Cake",
      "time": "45 min",
      "ingredients": [
        {
          "ingredient_name": "Flour",
          "ingredient_count": "2",
          "ingredient_unit": "mugs"
        },
        {
          "ingredient_name": "Strawberries",
          "ingredient_count": "8"
        },
        {
          "ingredient_name": "Coffee Beans",
          "ingredient_count": "2.5",
          "ingredient_unit": "tablespoons"
        },
        {
          "ingredient_name": "Cinnamon",
          "ingredient_count": "1"
        }
      ]
    },
    {
      "name": "Moonshine Muffin",
      "time": "30 min",
      "ingredients": [
        {
          "ingredient_name": "Brown sugar",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        },
        {
          "ingredient_name": "Blueberries",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        }
      ]
    }
  ]
}
```



Reading the file should be straightforward, so both these should work (files can be distinguished by an extension, for simplicity):

`~$ ./readDB -f original_database.xml`
`~$ ./readDB -f stolen_database.json`

Think of which kinds of objects are there in these databases and how they can be represented in code. Then, write an interface `DBReader` and two implementations of it - one for reading JSON and one for reading XML. Both of them should return the object of the same type as a result.

Make the code print JSON version of the database when it's reading from XML and vice versa. Both XML and JSON fields should be indented with 4 spaces ("pretty-printing").

<h2 id="ex01">Exercise 01: Assessing Damage</h2>

You've seen that the new database has modified versions of the same recipes, meaning there are several possible cases:

1) New cake is added or old one removed
2) Cooking time is different for the same cake
3) New ingredient is added or removed for the same cake. *Important:* the order of ingredients doesn't matter. Only the names are.
4) The count of units for the same ingredient has changed.
5) The unit itself for measuring the ingredient has changed.
6) Ingredient unit is missing or added

You may assume names are the same across both databases.

Your application should be runnable like this:

`~$ ./compareDB --old original_database.xml --new stolen_database.json`

It should work with both formats (JSON and XML) for original AND new database, reusing the code from Exercise 00.

The output should look like this (the same cases explained above):

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
```
<h2 id="ex02">Exercise 02: Dumps</h2>


Your program should take two filesystem dumps.

`~$ ./compareFS --old snapshot1.txt --new snapshot2.txt`

They are both plain text files, unsorted, and each of them includes a filepath on every like, like this:

```
/etc/stove/config.xml
/Users/baker/recipes/database.xml
/Users/baker/recipes/database_version3.yaml
/var/log/orders.log
/Users/baker/pokemon.avi
```

Your tool should output the very similar thing to a previous code (without CHANGED case though):

```
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
```

There is one issue though - the files can be really big, so you can assume both of them won't fit into RAM on the same time. There are two possible ways to overcome this - either to compress the file in memory somehow, or just read one of them and then avoid reading the other.
